alias gln='gleam new --skip-github --name main . && gleam add stdin'
alias ojt='gleam build -t javascript && oj t -c "gleam run -t javascript --no-print-progress" -d ./tests -N'
alias acs='esbuild build/dev/javascript/main/gleam.main.mjs --bundle --platform=node --target=node18.16 --outdir=dist --allow-overwrite --minify && acc s -- dist/gleam.main.js --language 5009'
