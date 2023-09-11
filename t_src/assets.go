package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets5e64cac0290a25853f1cc870097dd7f7af8bfe44 = "<!DOCTYPE html>\n<html lang=\"en\">\n\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n    <title>Embed Testing</title>\n</head>\n\n<body>\n    {{template \"content\" .}}\n</body>\n\n</html>"
var _Assets1c279dfbf695da97ff01d29135ca979613ca5da8 = "<!DOCTYPE html>\n<html lang=\"en\">\n\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n    <title>Empty</title>\n</head>\n\n<body>\n    <h1>Name: {{.Name}}</h1>\n</body>\n\n</html>"
var _Assetsbaaa8e8121df89b600c633bd5d246feb633b7468 = "{{define \"content\"}}\n<h1>About</h1>\n{{end}}"
var _Assetsb74dd11da69bf1cce8006f088e75315d093c2257 = "{{ define \"content\" }}\n<h1>{{.Name}}</h1>\n{{end}}"
var _Assetsd6ed3672ad756aaada55708774e589bef3ed43a1 = "package temp\n\nimport (\n\t\"time\"\n\n\t\"github.com/jessevdk/go-assets\"\n)\n\nvar _Assets1c279dfbf695da97ff01d29135ca979613ca5da8 = \"<!DOCTYPE html>\\n<html lang=\\\"en\\\">\\n\\n<head>\\n    <meta charset=\\\"UTF-8\\\">\\n    <meta name=\\\"viewport\\\" content=\\\"width=device-width, initial-scale=1.0\\\">\\n    <title>Empty</title>\\n</head>\\n\\n<body>\\n    <h1>{{.Name}}</h1>\\n</body>\\n\\n</html>\"\nvar _Assets5e64cac0290a25853f1cc870097dd7f7af8bfe44 = \"<!DOCTYPE html>\\n<html lang=\\\"en\\\">\\n\\n<head>\\n    <meta charset=\\\"UTF-8\\\">\\n    <meta name=\\\"viewport\\\" content=\\\"width=device-width, initial-scale=1.0\\\">\\n    <title>Embed Testing</title>\\n</head>\\n\\n<body>\\n    {{template \\\"content\\\" .}}\\n</body>\\n\\n</html>\"\nvar _Assetsb74dd11da69bf1cce8006f088e75315d093c2257 = \"{{ define \\\"content\\\" }}\\n<h1>{{.Name}}</h1>\\n{{end}}\"\nvar _Assetsbaaa8e8121df89b600c633bd5d246feb633b7468 = \"{{define \\\"content\\\"}}\\n<h1>About</h1>\\n{{end}}\"\nvar _Assets15713bfc125da850ab8dfe59ea7711f23e0b1ce8 = \"{{define \\\"content\\\"}}\\n<h1>Index</h1>\\n{{end}}\"\n\n// Assets returns go-assets FileSystem\nvar Assets = assets.NewFileSystem(map[string][]string{\"/\": []string{\"temp\"}, \"/temp\": []string{\"empty.html\", \"empty_content.html\"}, \"/temp/content\": []string{\"about.html\", \"index.html\"}, \"/temp/layout\": []string{\"base.html\"}}, map[string]*assets.File{\n\t\"/temp/empty.html\": &assets.File{\n\t\tPath:     \"/temp/empty.html\",\n\t\tFileMode: 0x1a4,\n\t\tMtime:    time.Unix(1694305870, 1694305870103098264),\n\t\tData:     []byte(_Assets1c279dfbf695da97ff01d29135ca979613ca5da8),\n\t}, \"/\": &assets.File{\n\t\tPath:     \"/\",\n\t\tFileMode: 0x800001ed,\n\t\tMtime:    time.Unix(1694212172, 1694212172331658722),\n\t\tData:     nil,\n\t}, \"/temp/layout\": &assets.File{\n\t\tPath:     \"/temp/layout\",\n\t\tFileMode: 0x800001ed,\n\t\tMtime:    time.Unix(1694209595, 1694209595375918147),\n\t\tData:     nil,\n\t}, \"/temp/empty_content.html\": &assets.File{\n\t\tPath:     \"/temp/empty_content.html\",\n\t\tFileMode: 0x1a4,\n\t\tMtime:    time.Unix(1694305918, 1694305918358361001),\n\t\tData:     []byte(_Assetsb74dd11da69bf1cce8006f088e75315d093c2257),\n\t}, \"/temp/content\": &assets.File{\n\t\tPath:     \"/temp/content\",\n\t\tFileMode: 0x800001ed,\n\t\tMtime:    time.Unix(1694209617, 1694209617615896031),\n\t\tData:     nil,\n\t}, \"/temp/content/about.html\": &assets.File{\n\t\tPath:     \"/temp/content/about.html\",\n\t\tFileMode: 0x1a4,\n\t\tMtime:    time.Unix(1694305881, 1694305881074930634),\n\t\tData:     []byte(_Assetsbaaa8e8121df89b600c633bd5d246feb633b7468),\n\t}, \"/temp/content/index.html\": &assets.File{\n\t\tPath:     \"/temp/content/index.html\",\n\t\tFileMode: 0x1a4,\n\t\tMtime:    time.Unix(1694305919, 1694305919862338022),\n\t\tData:     []byte(_Assets15713bfc125da850ab8dfe59ea7711f23e0b1ce8),\n\t}, \"/temp\": &assets.File{\n\t\tPath:     \"/temp\",\n\t\tFileMode: 0x800001ed,\n\t\tMtime:    time.Unix(1694305817, 1694305817595900490),\n\t\tData:     nil,\n\t}, \"/temp/layout/base.html\": &assets.File{\n\t\tPath:     \"/temp/layout/base.html\",\n\t\tFileMode: 0x1a4,\n\t\tMtime:    time.Unix(1694305918, 1694305918962351773),\n\t\tData:     []byte(_Assets5e64cac0290a25853f1cc870097dd7f7af8bfe44),\n\t}}, \"\")\n"
var _Assets15713bfc125da850ab8dfe59ea7711f23e0b1ce8 = "{{define \"content\"}}\n<h1>Index</h1>\n{{end}}"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"temp"}, "/temp": []string{"empty.html", "empty_content.html", "assets.go"}, "/temp/content": []string{"about.html", "index.html"}, "/temp/layout": []string{"base.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1694384204, 1694384204027784986),
		Data:     nil,
	}, "/temp/layout": &assets.File{
		Path:     "/temp/layout",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1694209595, 1694209595375918147),
		Data:     nil,
	}, "/temp/layout/base.html": &assets.File{
		Path:     "/temp/layout/base.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1694305918, 1694305918962351773),
		Data:     []byte(_Assets5e64cac0290a25853f1cc870097dd7f7af8bfe44),
	}, "/temp/empty.html": &assets.File{
		Path:     "/temp/empty.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1694385305, 1694385305384232988),
		Data:     []byte(_Assets1c279dfbf695da97ff01d29135ca979613ca5da8),
	}, "/temp/content/about.html": &assets.File{
		Path:     "/temp/content/about.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1694305881, 1694305881074930634),
		Data:     []byte(_Assetsbaaa8e8121df89b600c633bd5d246feb633b7468),
	}, "/temp": &assets.File{
		Path:     "/temp",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1694384204, 1694384204027784986),
		Data:     nil,
	}, "/temp/empty_content.html": &assets.File{
		Path:     "/temp/empty_content.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1694305918, 1694305918358361001),
		Data:     []byte(_Assetsb74dd11da69bf1cce8006f088e75315d093c2257),
	}, "/temp/assets.go": &assets.File{
		Path:     "/temp/assets.go",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1694384241, 1694384241365996977),
		Data:     []byte(_Assetsd6ed3672ad756aaada55708774e589bef3ed43a1),
	}, "/temp/content": &assets.File{
		Path:     "/temp/content",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1694209617, 1694209617615896031),
		Data:     nil,
	}, "/temp/content/index.html": &assets.File{
		Path:     "/temp/content/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1694305919, 1694305919862338022),
		Data:     []byte(_Assets15713bfc125da850ab8dfe59ea7711f23e0b1ce8),
	}}, "")
