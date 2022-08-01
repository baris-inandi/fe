echo "Creeating release $1"
git add -A
git commit --allow-empty -am "Release: $1"
gh release create $1 --generate-notes
git pull
git push
git tag --list
