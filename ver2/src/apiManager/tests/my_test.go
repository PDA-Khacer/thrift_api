package test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"managerStudent/ver2/src/apiManager/models"
	_ "managerStudent/ver2/src/apiManager/routers"
	"net/http"
	url2 "net/url"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestAddStudent(t *testing.T) {
	requestBody, err := json.Marshal(models.Student{
		Infor: &models.StudentInfor{ID: "99", Name: "abc"},
	})
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	reqBody := ioutil.NopCloser(strings.NewReader(string(requestBody)))
	url, err := url2.Parse("http://localhost:8080/v1/student/create")
	if err != nil {
		t.Error(err)
		return
	}
	rq := &http.Request{
		Method: "POST",
		URL:    url,
		Header: map[string][]string{
			"content-type": {"application/json; charset=UTF-8"},
		},
		Body: reqBody,
	}
	r, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	r.Body.Close()
	return
}

func TestAddClass(t *testing.T) {
	requestBody, err := json.Marshal(models.ClassC{
		Infor: &models.ClassInfor{ID: "09", Name: "abc"},
	})
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	reqBody := ioutil.NopCloser(strings.NewReader(string(requestBody)))
	url, err := url2.Parse("http://localhost:8080/v1/classC/create")
	if err != nil {
		t.Error(err)
		return
	}
	rq := &http.Request{
		Method: "POST",
		URL:    url,
		Header: map[string][]string{
			"content-type": {"application/json; charset=UTF-8"},
		},
		Body: reqBody,
	}
	r, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	r.Body.Close()
	return
}

func TestGetStudent(t *testing.T) {
	url, err := url2.Parse("http://localhost:8080/v1/student/99")
	if err != nil {
		t.Error(err)
		return
	}
	rq := &http.Request{
		Method: "GET",
		URL:    url,
		Header: map[string][]string{
			"content-type": {"application/json; charset=UTF-8"},
		},
	}
	r, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	r.Body.Close()
	return
}

func TestGetClass(t *testing.T) {
	url, err := url2.Parse("http://localhost:8080/v1/classC/09")
	if err != nil {
		t.Error(err)
		return
	}
	rq := &http.Request{
		Method: "GET",
		URL:    url,
		Header: map[string][]string{
			"content-type": {"application/json; charset=UTF-8"},
		},
	}
	r, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	r.Body.Close()
	return
}

func TestAllClass(t *testing.T) {
	url, err := url2.Parse("http://localhost:8080/v1/classC/all")
	if err != nil {
		t.Error(err)
		return
	}
	rq := &http.Request{
		Method: "GET",
		URL:    url,
		Header: map[string][]string{
			"content-type": {"application/json; charset=UTF-8"},
		},
	}
	r, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	r.Body.Close()
	return
}

func TestAllStudent(t *testing.T) {
	url, err := url2.Parse("http://localhost:8080/v1/student/all")
	if err != nil {
		t.Error(err)
		return
	}
	rq := &http.Request{
		Method: "GET",
		URL:    url,
		Header: map[string][]string{
			"content-type": {"application/json; charset=UTF-8"},
		},
	}
	r, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	r.Body.Close()
	return
}

func TestUpdateStudent(t *testing.T) {
	requestBody, err := json.Marshal(models.Student{
		Infor: &models.StudentInfor{ID: "99", Name: "hello"},
	})
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	reqBody := ioutil.NopCloser(strings.NewReader(string(requestBody)))
	url, err := url2.Parse("http://localhost:8080/v1/student/")
	if err != nil {
		t.Error(err)
		return
	}
	rq := &http.Request{
		Method: "PUT",
		URL:    url,
		Header: map[string][]string{
			"content-type": {"application/json; charset=UTF-8"},
		},
		Body: reqBody,
	}
	r, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	r.Body.Close()
	return
}

func TestUpdateClass(t *testing.T) {
	requestBody, err := json.Marshal(models.ClassC{
		Infor: &models.ClassInfor{ID: "09", Name: "tesst"},
	})
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	reqBody := ioutil.NopCloser(strings.NewReader(string(requestBody)))
	url, err := url2.Parse("http://localhost:8080/v1/classC/")
	if err != nil {
		t.Error(err)
		return
	}
	rq := &http.Request{
		Method: "PUT",
		URL:    url,
		Header: map[string][]string{
			"content-type": {"application/json; charset=UTF-8"},
		},
		Body: reqBody,
	}
	r, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	r.Body.Close()
	return
}

func TestDeleteStudent(t *testing.T) {
	url, err := url2.Parse("http://localhost:8080/v1/student/99")
	if err != nil {
		t.Error(err)
		return
	}
	rq := &http.Request{
		Method: "DELETE",
		URL:    url,
		Header: map[string][]string{
			"content-type": {"application/json; charset=UTF-8"},
		},
	}
	r, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	r.Body.Close()
	return
}

func TestDeleteClass(t *testing.T) {
	url, err := url2.Parse("http://localhost:8080/v1/classC/09")
	if err != nil {
		t.Error(err)
		return
	}
	rq := &http.Request{
		Method: "DELETE",
		URL:    url,
		Header: map[string][]string{
			"content-type": {"application/json; charset=UTF-8"},
		},
	}
	r, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	r.Body.Close()
	return
}

func TestAddToClass(t *testing.T) {
	requestBody, err := json.Marshal(models.StudentClassC{
		ClassId:   "09",
		StudentId: "99",
	})
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	reqBody := ioutil.NopCloser(strings.NewReader(string(requestBody)))
	url, err := url2.Parse("http://localhost:8080/v1/student/AddToClass")
	if err != nil {
		t.Error(err)
		return
	}
	rq := &http.Request{
		Method: "PUT",
		URL:    url,
		Header: map[string][]string{
			"content-type": {"application/json; charset=UTF-8"},
		},
		Body: reqBody,
	}
	r, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	r.Body.Close()
	return
}

func TestOutClass(t *testing.T) {
	requestBody, err := json.Marshal(models.StudentClassC{
		ClassId:   "09",
		StudentId: "99",
	})
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	reqBody := ioutil.NopCloser(strings.NewReader(string(requestBody)))
	url, err := url2.Parse("http://localhost:8080/v1/student/OutClass")
	if err != nil {
		t.Error(err)
		return
	}
	rq := &http.Request{
		Method: "DELETE",
		URL:    url,
		Header: map[string][]string{
			"content-type": {"application/json; charset=UTF-8"},
		},
		Body: reqBody,
	}
	r, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	_, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	r.Body.Close()
	return
}
