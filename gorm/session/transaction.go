package session

import "gorm/log"

// 这里输出必须加上err，不然err会报错，因为输出的是s.tx，所以err也需要先定义出来
func (s *Session) Begin() (err error) {
	log.Info("transaction begin")
	//var err error
	if s.tx, err = s.db.Begin(); err != nil {
		log.Error(err)
		return
	}
	return
}

func (s *Session) Commit() (err error) {
	log.Info("transaction commit")
	//var err error
	if err = s.tx.Commit(); err != nil {
		log.Error(err)
		return
	}
	return
}

func (s *Session) Rollback() (err error) {
	log.Info("transaction rollback")
	//var err error
	if err = s.tx.Rollback(); err != nil {
		log.Error(err)
		return
	}
	return
}
