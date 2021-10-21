package json

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type codeResponse struct {
	Tree     *codeNode `json:"tree"`
	Username string    `json:"username"`
}

type codeNode struct {
	Name     string      `json:"name"`
	Kids     []*codeNode `json:"kids"`
	CLWeight float64     `json:"cl_weight"`
	Touches  int         `json:"touches"`
	MinT     int64       `json:"min_t"`
	MaxT     int64       `json:"max_t"`
	MeanT    int64       `json:"mean_t"`
}

var codeJSON []byte
var codeStruct codeResponse

func codeInit() {
	f, err := os.Open("testdata/code.json.gz")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(gz)
	if err != nil {
		panic(err)
	}

	codeJSON = data

	if err := Unmarshal(codeJSON, &codeStruct); err != nil {
		panic("unmarshal code.json: " + err.Error())
	}

	if data, err = Marshal(&codeStruct); err != nil {
		panic("marshal code.json: " + err.Error())
	}

	if !bytes.Equal(data, codeJSON) {
		println("different lengths", len(data), len(codeJSON))
		for i := 0; i < len(data) && i < len(codeJSON); i++ {
			if data[i] != codeJSON[i] {
				println("re-marshal: changed at byte", i)
				println("orig: ", string(codeJSON[i-10:i+10]))
				println("new: ", string(data[i-10:i+10]))
				break
			}
		}
		panic("re-marshal code.json: different result")
	}
}

func BenchmarkCodeEncoder(b *testing.B) {
	b.ReportAllocs()
	if codeJSON == nil {
		b.StopTimer()
		codeInit()
		b.StartTimer()
	}
	b.RunParallel(func(pb *testing.PB) {
		enc := NewEncoder(ioutil.Discard)
		for pb.Next() {
			if err := enc.Encode(&codeStruct); err != nil {
				b.Fatal("Encode:", err)
			}
		}
	})
	b.SetBytes(int64(len(codeJSON)))
}

func BenchmarkCodeMarshal(b *testing.B) {
	b.ReportAllocs()
	if codeJSON == nil {
		b.StopTimer()
		codeInit()
		b.StartTimer()
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := Marshal(&codeStruct); err != nil {
				b.Fatal("Marshal:", err)
			}
		}
	})
	b.SetBytes(int64(len(codeJSON)))
}

func benchMarshalBytes(n int) func(*testing.B) {
	sample := []byte("hello world")
	// Use a struct pointer, to avoid an allocation when passing it as an
	// interface parameter to Marshal.
	v := &struct {
		Bytes []byte
	}{
		bytes.Repeat(sample, (n/len(sample))+1)[:n],
	}
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, err := Marshal(v); err != nil {
				b.Fatal("Marshal:", err)
			}
		}
	}
}

func BenchmarkMarshalBytes(b *testing.B) {
	b.ReportAllocs()
	// 32 fits within encodeState.scratch.
	b.Run("32", benchMarshalBytes(32))
	// 256 doesn't fit in encodeState.scratch, but is small enough to
	// allocate and avoid the slower base64.NewEncoder.
	b.Run("256", benchMarshalBytes(256))
	// 4096 is large enough that we want to avoid allocating for it.
	b.Run("4096", benchMarshalBytes(4096))
}

func BenchmarkCodeDecoder(b *testing.B) {
	b.ReportAllocs()
	if codeJSON == nil {
		b.StopTimer()
		codeInit()
		b.StartTimer()
	}
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		dec := NewDecoder(&buf)
		var r codeResponse
		for pb.Next() {
			buf.Write(codeJSON)
			// hide EOF
			buf.WriteByte('\n')
			buf.WriteByte('\n')
			buf.WriteByte('\n')
			if err := dec.Decode(&r); err != nil {
				b.Fatal("Decode:", err)
			}
		}
	})
	b.SetBytes(int64(len(codeJSON)))
}

func BenchmarkUnicodeDecoder(b *testing.B) {
	b.ReportAllocs()
	j := []byte(`"\uD83D\uDE01"`)
	b.SetBytes(int64(len(j)))
	r := bytes.NewReader(j)
	dec := NewDecoder(r)
	var out string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := dec.Decode(&out); err != nil {
			b.Fatal("Decode:", err)
		}
		i64, e := r.Seek(0, 0)
		b.Logf("%#v,%#v\n", i64, e)
	}
}

func BenchmarkDecoderStream(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	var buf bytes.Buffer
	dec := NewDecoder(&buf)
	buf.WriteString(`"` + strings.Repeat("x", 1000000) + `"` + "\n\n\n")
	var x interface{}
	if err := dec.Decode(&x); err != nil {
		b.Fatal("Decode:", err)
	}
	ones := strings.Repeat(" 1\n", 300000) + "\n\n\n"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if i%300000 == 0 {
			buf.WriteString(ones)
		}
		x = nil
		if err := dec.Decode(&x); err != nil || x != 1.0 {
			b.Fatalf("Decode: %v after %d", err, i)
		}
	}
}

func BenchmarkCodeUnmarshal(b *testing.B) {
	b.ReportAllocs()
	if codeJSON == nil {
		b.StopTimer()
		codeInit()
		b.StartTimer()
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r codeResponse
			if err := Unmarshal(codeJSON, &r); err != nil {
				b.Fatal("Unmarshal:", err)
			}
		}
	})
	b.SetBytes(int64(len(codeJSON)))
}

func BenchmarkCodeUnmarshalReuse(b *testing.B) {
	b.ReportAllocs()
	if codeJSON == nil {
		b.StopTimer()
		codeInit()
		b.StartTimer()
	}
	b.RunParallel(func(pb *testing.PB) {
		var r codeResponse
		for pb.Next() {
			if err := Unmarshal(codeJSON, &r); err != nil {
				b.Fatal("Unmarshal:", err)
			}
		}
	})
	b.SetBytes(int64(len(codeJSON)))
}

func BenchmarkUnmarshalString(b *testing.B) {
	b.ReportAllocs()
	data := []byte(`"hello, world"`)
	b.RunParallel(func(pb *testing.PB) {
		var s string
		for pb.Next() {
			if err := Unmarshal(data, &s); err != nil {
				b.Fatal("Unmarshal:", err)
			}
		}
	})
}

func BenchmarkUnmarshalFloat64(b *testing.B) {
	b.ReportAllocs()
	data := []byte(`3.14`)
	b.RunParallel(func(pb *testing.PB) {
		var f float64
		for pb.Next() {
			if err := Unmarshal(data, &f); err != nil {
				b.Fatal("Unmarshal:", err)
			}
		}
	})
}

func BenchmarkUnmarshalInt64(b *testing.B) {
	b.ReportAllocs()
	data := []byte(`3`)
	b.RunParallel(func(pb *testing.PB) {
		var x int64
		for pb.Next() {
			if err := Unmarshal(data, &x); err != nil {
				b.Fatal("Unmarshal:", err)
			}
		}
	})
}

func BenchmarkIssue10335(b *testing.B) {
	b.ReportAllocs()
	j := []byte(`{"a":{ }}`)
	b.RunParallel(func(pb *testing.PB) {
		var s struct{}
		for pb.Next() {
			if err := Unmarshal(j, &s); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkIssue34127(b *testing.B) {
	b.ReportAllocs()
	j := struct {
		Bar string `json:"bar,string"`
	}{
		Bar: `foobar`,
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := Marshal(&j); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkUnmapped(b *testing.B) {
	b.ReportAllocs()
	j := []byte(`{"s": "hello", "y": 2, "o": {"x": 0}, "a": [1, 99, {"x": 1}]}`)
	b.RunParallel(func(pb *testing.PB) {
		var s struct{}
		for pb.Next() {
			if err := Unmarshal(j, &s); err != nil {
				b.Fatal(err)
			}
		}
	})
}

type TS struct {
	A map[string]interface{}
	B []string
	C int
	D string
}

func TestToJsonStringPretty(t *testing.T) {
	t.Logf("%s\n", ToJSONStringPretty(TS{
		map[string]interface{}{
			"a": "a",
			"b": 1,
			"c": 1.1,
		},
		[]string{
			"1", "2",
		},
		1,
		"2",
	}))
}

func TestFromJsonString(t *testing.T) {
	ts := &TS{}
	t.Logf("%s\n", FromJSONString(`{
            "A": {
                "a": "a",
                "b": 1,
                "c": 1.1
            },
            "B": [
                "1",
                "2"
            ],
            "C": 1,
            "D": "2"
        }`, ts))

	t.Logf("%#v\n", ts)
}

func TestJSONNumber(t *testing.T) {

	type SubjectRule struct {
		Weight   float64 `json:"weight"`
		Operator string  `json:"operator"`
		MinValue float64 `json:"minValue"`
		MaxValue float64 `json:"maxValue"`
		Summary  string  `json:"summary"`
		Callback string  `json:"callback"`
	}

	s := `
    {
        "weight":0,
        "operator":">",
        "minValue":30,
        "maxValue":0,
        "summary":">30分钟得分",
        "callback":"syncNewScore"
	}
`
	sr := &SubjectRule{}

	er := Unmarshal([]byte(s), sr)
	t.Logf("%#v\n", er)
	t.Logf("%#v\n", sr)

}
