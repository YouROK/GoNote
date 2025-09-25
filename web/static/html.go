package template

import (
	_ "embed"
)

//go:embed files/apple-touch-icon.png
var Filesappletouchiconpng []byte

//go:embed files/banner.png
var Filesbannerpng []byte

//go:embed files/css/gonote.css
var Filescssgonotecss []byte

//go:embed files/favicon-96x96.png
var Filesfavicon96x96png []byte

//go:embed files/favicon.ico
var Filesfaviconico []byte

//go:embed files/favicon.svg
var Filesfaviconsvg []byte

//go:embed files/js/editor.js
var Filesjseditorjs []byte

//go:embed files/js/shared.js
var Filesjssharedjs []byte

//go:embed files/js/viewer.js
var Filesjsviewerjs []byte

//go:embed files/robots.txt
var Filesrobotstxt []byte

//go:embed files/site.webmanifest
var Filessitewebmanifest []byte

//go:embed files/web-app-manifest-192x192.png
var Fileswebappmanifest192x192png []byte

//go:embed files/web-app-manifest-512x512.png
var Fileswebappmanifest512x512png []byte
