package sequence

import (
	"short_url/base"
	"short_url/base/tool"
)

func (s *service) GetBorrowOrder() (int64, error) {
	var (
		err  error
		data int64
	)
	if err = s.r.Incr(base.REDIS_SEQUENCE_GENERATOR).Err(); err != nil {
		tool.GetLogger().Error("[GetBorrowOrder] s.r.Incr " + base.REDIS_SEQUENCE_GENERATOR)
		return 0, err
	}
	if data, err = s.r.Get(base.REDIS_SEQUENCE_GENERATOR).Int64(); err != nil {
		tool.GetLogger().Error("[GetBorrowOrder] s.r.Get " + base.REDIS_SEQUENCE_GENERATOR)
		return 0, err
	}
	return data, nil
}
