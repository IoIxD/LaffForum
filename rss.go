package main

import (
	"fmt"
	"net/http"
)


func RSSServe(w http.ResponseWriter, r *http.Request, values []string) {
	if(len(values) <= 1) {
		return
	}
	w.Header().Set("Content-Type", "text/xml")
	w.Header().Set("Content-Name", values[1]+".xml")
	w.Write([]byte(`<?xml version="1.0" encoding="UTF-8" ?>
		<rss version="2.0">
		<channel>
			<link>https://`+r.Host+`</link>
		`))

	switch(values[1]) {
		case "topic":
			fmt.Println(len(values))
			if(len(values) <= 2) {
				w.Write([]byte(`<title>Must specify topic name</title>`))
				break
			}
			w.Write([]byte(`<title>`+Capitalize(values[2])+`</title>`))
			w.Write([]byte(`<description> Posts in `+Capitalize(values[2])+`</description>`))
			result := GetPostsBySectionName(values[2])
			if(result.Error != nil) {
				w.Write([]byte(`<item> 
									<title>Error</title>
									<description>`+result.Error.Error()+`</description>
								</item>`))
				break
			}
			for _, v := range result.Posts {
				w.Write([]byte(`
					<item>
						<title>`+v.Subject+`</title>
						<description>`+TrimForMeta(v.Contents)+`</description>
						<link>https://`+r.Host+`/post/`+fmt.Sprint(v.ID)+`</link>
					</item>
					`))
			}
		case "post":
			// Try and get the post information.
			result := GetPostInfo(values[2])
			if(result.Error != nil) {
				w.Write([]byte(`<item> 
									<title>Error</title>
									<description>`+result.Error.Error()+`</description>
								</item>`))
				break
			}

			w.Write([]byte(`<title>`+result.Subject+`</title>`))
			w.Write([]byte(`<description>"`+result.Subject+`" and its replies</description>`))

			// Show the original post as the first result.
			w.Write(XMLShowPost(`https://`+r.Host, result))

			replies := GetPostsInReplyTo(result.ID)
			if(replies.Error != nil) {
				w.Write([]byte(`<item> 
									<title>Error</title>
									<description>`+replies.Error.Error()+`</description>
								</item>`))
				break
			}
			for _, v := range replies.Posts {
				w.Write(XMLShowPost(`https://`+r.Host, v))
			}
	}
	w.Write([]byte(`
			</channel>
		</rss>`))
}

func XMLShowPost(url string, post Post) ([]byte) {
	author := GetUsernameByID(post.Author)
	var authorname string
	if(author.Error != nil) {
		authorname = author.Error.Error()
	} else {
		authorname = author.Result.(string)
	}
	return []byte(`
		<item>
			<title>`+authorname+`: "`+post.Subject+`"</title>
			<description>`+TrimForMeta(post.Contents)+`</description>
			<link>`+url+`/post/`+fmt.Sprint(post.ID)+`</link>
		</item>
	`)
}