git init
git remote add origin https://github.com
git branch -M main
git add .
git commit -m "Official Launch Release"
git tag v0.1.0
git push -u origin main --tags
$env:GOPROXY="https://golang.org"
go list -m ://github.com
