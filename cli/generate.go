package cli

//go:generate sh -c "command -v gosubc >/dev/null 2>&1 && gosubc generate --dir .. || go run github.com/arran4/go-subcommand/cmd/gosubc generate --dir .."
//go:generate sed -i "s/\\tGallery()/\\tcli.Gallery()/" ../cmd/frames/gallery.go
//go:generate sed -i "s/\\tGenerate()/\\tcli.Generate()/" ../cmd/frames/generate.go
//go:generate sed -i "s/\\tGenWood()/\\tcli.GenWood()/" ../cmd/frames/wood.go
