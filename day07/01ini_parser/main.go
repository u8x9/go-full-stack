package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type PostgreConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	DbName   string `ini:"dbname"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

func InitConfig(filename string, cfg interface{}) error {
	v := reflect.ValueOf(cfg)
	if v.Kind() != reflect.Ptr {
		return errors.New("需要一个指针")
	}
	if v.Elem().Kind() != reflect.Struct {
		return errors.New("需要一个struct指针")
	}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	r := bufio.NewReader(file)
    section:=""
    structName:=strings.ToLower(v.Elem().Type().Name())
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		line = strings.TrimSpace(line)
        if len(line) == 0 {
            continue
        }
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") { // 跳过注释
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") { // 节
            section = strings.TrimLeft(line,"[")
            section = strings.TrimRight(section,"]")
            section = strings.ToLower(section)
			//fmt.Println(v.Elem().Type().Name())
            continue
		}
        if !strings.HasPrefix(structName, section) {
            continue
        }
        symbolIndex := strings.Index(line,"=")
        if symbolIndex == -1 {
            continue
        }
        //fmt.Println(section, line)
        cfgKey := line[:symbolIndex]
        cfgVal := line[symbolIndex+1:]
        //fmt.Println(cfgKey, cfgVal)
        for i:=0; i < v.Elem().Type().NumField(); i++{
            field := v.Elem().Type().Field(i)
            if field.Tag.Get("ini") != cfgKey {
                continue
            }
            switch field.Type.Kind() {
            case reflect.String:
                v.Elem().FieldByName(field.Name).SetString(cfgVal)
            case reflect.Int:
                vv,err := strconv.ParseInt(cfgVal, 10, 32)
                if err != nil{
                    continue
                }
                v.Elem().FieldByName(field.Name).SetInt(vv)
            }
        }
	}
	return nil
}

func main() {
	var pc PostgreConfig
	err := InitConfig("config.ini", &pc)
	if err != nil {
		fmt.Println(err)
		return
	}
    fmt.Printf("%#v\n", pc)
	var rc RedisConfig
	err = InitConfig("config.ini", &rc)
	if err != nil {
		fmt.Println(err)
		return
	}
    fmt.Printf("%#v\n", rc)
}
