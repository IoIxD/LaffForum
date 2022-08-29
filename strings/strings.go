package strings

import (
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/IoIxD/LaffForum/database"
	"github.com/gomarkdown/markdown"
)

// Capitalize a string
func Capitalize(value string) string {
	// Treat dashes as spaces
	value = strings.Replace(value, "-", " ", 99)
	valuesplit := strings.Split(value, " ")
	var result string
	for _, v := range valuesplit {
		if(len(v) <= 0) {
			continue
		}
		result += strings.ToUpper(v[:1])
		result += v[1:] + " "
	}
	return result
}

// Trim a string to 128 characters, for meta tags.
func TrimForMeta(value string) string {
	if(len(value) <= 127) {
		return value
	}
	return value[:128]+"..."
}

// Print the server date three months from now 
func PrintThreeMonthsFromNow() string {
	future := time.Now().Add(time.Hour*2190)
	return future.Format("Jan 02 2006, 03:04:05PM -0700")
}

// Parsing a markdown string.

func Markdown(val string) (string) {
	val = template.HTMLEscapeString(val)
	val = strings.Replace(val,"{{QUAKE}}",
		`<a href='/WebQuake/Client/index.htm'>
			<iframe width='1024' height='768' src='/WebQuake/Client/index.htm' class='quake-iframe'></iframe>
		</a>`,1)
	return string(markdown.ToHTML([]byte(val),nil,nil))
}

func HTMLEscape(val string) (string) {
	return template.HTMLEscapeString(val)
}

// Function for formatting a timestamp as "x hours ago"
func PrettyTime(unixTime int) (result database.GenericResult) {
	unixTimeDur, err := time.ParseDuration(fmt.Sprintf("%vs", time.Now().Unix()-int64(unixTime)))
	if err != nil {
		result.Error = err
		return
	}

	if unixTimeDur.Hours() >= 8760 {
		result.Result = fmt.Sprintf("%0.f years ago", unixTimeDur.Hours()/8760)
		return
	}
	if unixTimeDur.Hours() >= 730 {
		result.Result = fmt.Sprintf("%0.f months ago", unixTimeDur.Hours()/730)
		return
	}
	if unixTimeDur.Hours() >= 168 {
		result.Result = fmt.Sprintf("%0.f weeks ago", unixTimeDur.Hours()/168)
		return
	}
	if unixTimeDur.Hours() >= 24 {
		result.Result = fmt.Sprintf("%0.f days ago", unixTimeDur.Hours()/24)
		return
	}
	if unixTimeDur.Hours() >= 1 {
		result.Result = fmt.Sprintf("%0.f hours ago", unixTimeDur.Hours())
		return
	}
	if unixTimeDur.Minutes() >= 1 {
		result.Result = fmt.Sprintf("%0.f minutes ago", unixTimeDur.Minutes())
		return
	}
	result.Result = fmt.Sprintf("%0.f seconds ago", unixTimeDur.Seconds())
	return
}