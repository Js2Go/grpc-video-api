package dao

import (
	"video/backend/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (d *Dao) GetVideoInfo(id string) *model.Video {
	querySQL, err := d.DB.Prepare("SELECT title, note, pic_path, video_path FROM video WHERE id = ?")
	if err != nil {
		logrus.Errorf("Prepare Select SQL Error: %s", err.Error())
		return nil
	}

	defer querySQL.Close()
	v := new(model.Video)
	err = querySQL.QueryRow(id).Scan(&v.Title, &v.Note, &v.PicPath, &v.VideoPath)
	if err != nil {
		logrus.Errorf("Query Video Error: %s", err.Error())
	}
	return v
}

func (d *Dao) CreateNewVideo(v *model.Video) {
	id := UUID()
	title := v.Title
	note := v.Note
	picPath := v.PicPath
	videoPath := v.VideoPath

	insSQL, err := d.DB.Prepare("INSERT INTO video (id, title, note, pic_path, video_path) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		logrus.Errorf("Prepare Insert SQL Error: %s", err.Error())
	}

	_, err = insSQL.Exec(id, title, note, picPath, videoPath)
	if err != nil {
		logrus.Errorf("Insert Video Info SQL Error: %s", err.Error())
	}
	defer insSQL.Close()
	return
}

func UUID() string {
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Info("UUID Error: %s", err.Error())
	}
	return id.String()
}
