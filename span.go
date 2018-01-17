package zerolog

import (
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"encoding/json"
)

type span struct {
	context    []log.Field
	originSpan opentracing.Span
}

func newSpan(s opentracing.Span) *span {
	return &span{
		context:    make([]log.Field, 0),
		originSpan: s,
	}
}

func (s *span) logWithFields(msg string, level Level, fields ...log.Field) {
	fields = append(fields, makeSpanInfoLogFields(msg, level)...)
	fields = append(fields, s.context...)
	s.originSpan.LogFields(fields...)
}

func makeSpanInfoLogFields(msg string, level Level) []log.Field {
	return []log.Field{
		log.String("level", level.String()),
		log.String("event", msg),
	}
}

//TODO: Not dereference twice
type spanFields []log.Field

func newSpanFields() spanFields {
	return make([]log.Field, 0, 2)
}

func (sf *spanFields) Fields(fields map[string]interface{}) {
	if sf == nil {
		return
	}
}

func (sf *spanFields) Dict(key string, dict *Event) {
	if sf == nil {
		return
	}
	*sf = append(*sf, dict.spanFields...)
}

//TODO: implement marshaller
func (sf *spanFields) Array(key string, arr string) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.String(key, arr))
}

//TODO: implement marshaller
func (sf *spanFields) Object(key string, obj LogObjectMarshaler) {
	if sf == nil {
		return
	}
}

// Str adds the field key with val as a string to thf spanFields context.
func (sf *spanFields) Str(key, val string) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.String(key, val))
}

func (sf *spanFields) Strs(key string, vals []string) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.Object(key, vals))

}

func (sf *spanFields) Bytes(key string, val []byte) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.Object(key, val))
}

func (sf *spanFields) AnErr(key string, err error) {
	if sf == nil {
		return
	}
	if err == nil {
		return
	}

	*sf = append(*sf, log.Object(key, err))
}

func (sf *spanFields) Errs(key string, errs []error) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Object(key, errs))
}

func (sf *spanFields) Err(err error) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Error(err))
}

// Bool adds the field key with val as a bool to thf spanFields context.
func (sf *spanFields) Bool(key string, b bool) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Bool(key, b))
}

func (sf *spanFields) Bools(key string, b []bool) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Object(key, b))
}

func (sf *spanFields) Int(key string, i int) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Int(key, i))
}

func (sf *spanFields) Ints(key string, i []int) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Object(key, i))
}

//Int8 not support, casting to int
func (sf *spanFields) Int8(key string, i int8) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Int(key, int(i)))
}

//Int8 not support, casting to int
func (sf *spanFields) Ints8(key string, i []int8) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Object(key, i))
}

//int16 not support, casting to int
func (sf *spanFields) Int16(key string, i int16) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Int(key, int(i)))
}

func (sf *spanFields) Ints16(key string, i []int16) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Object(key, i))
}

func (sf *spanFields) Int32(key string, i int32) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Int32(key, i))
}

func (sf *spanFields) Ints32(key string, i []int32) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Object(key, i))
}

func (sf *spanFields) Int64(key string, i int64) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Int64(key, i))
}

func (sf *spanFields) Ints64(key string, i []int64) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.Object(key, i))
}

func (sf *spanFields) Uint(key string, i uint) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Uint32(key, uint32(i)))
}

func (sf *spanFields) Uints(key string, i []uint) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.Object(key, i))
}

func (sf *spanFields) Uint8(key string, i uint8) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Uint32(key, uint32(i)))
}

func (sf *spanFields) Uints8(key string, i []uint8) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.Object(key, i))
}

func (sf *spanFields) Uint16(key string, i uint16) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.Uint32(key, uint32(i)))
}

func (sf *spanFields) Uints16(key string, i []uint16) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.Object(key, i))
}

func (sf *spanFields) Uint32(key string, i uint32) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.Uint32(key, i))
}

func (sf *spanFields) Uints32(key string, i []uint32) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.Object(key, i))
}

func (sf *spanFields) Uint64(key string, i uint64) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.Uint64(key, i))
}

func (sf *spanFields) Uints64(key string, i []uint64) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.Object(key, i))
}

func (sf *spanFields) Float32(key string, f float32) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.Float32(key, f))
}

func (sf *spanFields) Floats32(key string, f []float32) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.Object(key, f))

}

func (sf *spanFields) Float64(key string, f float64) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.Float64(key, f))
}

func (sf *spanFields) Floats64(key string, f []float64) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.Object(key, f))
}

func (sf *spanFields) Timestamp() {
	return
}

func (sf *spanFields) Time(key string, t time.Time) {
	if sf == nil {
		return
	}
	*sf = append(*sf, log.String(key, t.Format(TimeFieldFormat)))
}

func (sf *spanFields) Times(key string, t []time.Time) {
	if sf == nil {
		return
	}

	times := make([]string, 0, len(t))
	for _, rawTime := range t {
		times = append(times, rawTime.Format(TimeFieldFormat))
	}

	*sf = append(*sf, log.Object(key, times))
}

func (sf *spanFields) Dur(key string, d time.Duration) {
	if sf == nil {
		return
	}

	*sf = append(*sf, log.Float64(key, float64(d)/float64(DurationFieldUnit)))
}

func (sf *spanFields) Durs(key string, d []time.Duration) {
	if sf == nil {
		return
	}

	durs := make([]float64, 0, len(d))
	for _, dur := range d {
		durs = append(durs, float64(dur)/float64(DurationFieldUnit))
	}

	*sf = append(*sf, log.Object(key, durs))
}

func (sf *spanFields) TimeDiff(key string, t time.Time, start time.Time) {
	if sf == nil {
		return
	}
	var d time.Duration
	if t.After(start) {
		d = t.Sub(start)
	}

	*sf = append(*sf, log.Float64(key, float64(d)/float64(DurationFieldUnit)))
}

// Interface adds the field key with i marshaled using reflection.
func (sf *spanFields) Interface(key string, i interface{}) {
	if sf == nil {
		return
	}

	marshaled, err := json.Marshal(i)
	if err != nil {
		//TODO: No panic
		panic(err)
	}

	*sf = append(*sf, log.String(key, string(marshaled)))
}
