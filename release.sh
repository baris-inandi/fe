echo "Creeating release $1"
git add -A
git commit --allow-empty -am "Release: $1"
git pull
git push
gh release create $1 --generate-notes
git tag --list
