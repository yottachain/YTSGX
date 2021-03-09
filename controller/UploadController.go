package controller

//UploadFile 文件上传
// func UploadFile(c *gin.Context) {

// 	var content_length int64
// 	content_length = c.Request.ContentLength
// 	if content_length <= 0 || content_length > 1024*1024*1024*2 {
// 		log.Printf("content_length error\n")
// 		return
// 	}
// 	content_type_, has_key := c.Request.Header["Content-Type"]
// 	if !has_key {
// 		log.Printf("Content-Type error\n")
// 		return
// 	}
// 	if len(content_type_) != 1 {
// 		log.Printf("Content-Type count error\n")
// 		return
// 	}
// 	content_type := content_type_[0]
// 	const BOUNDARY string = "; boundary="
// 	loc := strings.Index(content_type, BOUNDARY)
// 	if -1 == loc {
// 		log.Printf("Content-Type error, no boundary\n")
// 		return
// 	}
// 	boundary := []byte(content_type[(loc + len(BOUNDARY)):])
// 	log.Printf("[%s]\n\n", boundary)
// 	//
// 	read_data := make([]byte, 1024*12)
// 	var read_total int = 0
// 	for {
// 		file_header, file_data, err := s3server.ParseFromHead(read_data, read_total, append(boundary, []byte("\r\n")...), c.Request.Body)
// 		if err != nil {
// 			log.Printf("%v", err)
// 			return
// 		}
// 		log.Printf("file :%s\n", file_header.FileName)
// 		//
// 		f, err := os.Create(file_header.FileName)
// 		if err != nil {
// 			log.Printf("create file fail:%v\n", err)
// 			return
// 		}
// 		f.Write(file_data)
// 		file_data = nil

// 		//需要反复搜索boundary
// 		temp_data, reach_end, err := s3server.ReadToBoundary(boundary, c.Request.Body, f)
// 		f.Close()
// 		if err != nil {
// 			log.Printf("%v\n", err)
// 			return
// 		}
// 		if reach_end {
// 			break
// 		} else {
// 			copy(read_data[0:], temp_data)
// 			read_total = len(temp_data)
// 			continue
// 		}
// 	}
// 	//
// 	c.JSON(200, gin.H{
// 		"message": fmt.Sprintf("%s", "ok"),
// 	})

// }
