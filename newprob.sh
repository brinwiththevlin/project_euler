#!/usr/bin/env bash
# Scaffold a new Project Euler problem package and register it.
#
#   ./newprob.sh 7          -> creates p007/{p007.go,p007_test.go,run.go}
#   ./newprob.sh --proof 7  -> same, plus p007/p007_proof.tex
#                              (if p007 already exists, only adds the proof)
#   ./newprob.sh --sync     -> only regenerate problems.go and check registrations
#
# problems.go is always regenerated from the pNNN directories on disk, and every
# run.go is checked to make sure it registers the number its directory claims.
set -euo pipefail

cd "$(dirname "$0")"
module=$(awk '/^module /{print $2}' go.mod)

usage() { echo "usage: $0 [--proof] <problem-number> | --sync" >&2; exit 1; }

proof=0
if [[ ${1-} == --proof ]]; then
	proof=1
	shift
fi
[[ $# -eq 1 ]] || usage
[[ $proof -eq 0 || $1 != --sync ]] || usage

# fetch_title <n>: print the problem's title from projecteuler.net, LaTeX-escaped.
fetch_title() {
	curl -fsS --max-time 10 "https://projecteuler.net/problem=$1" |
		grep -oP '<h2>\K[^<]+' | head -n1 |
		sed -E -e 's/&amp;/\&/g' -e 's/([&%#_])/\\\1/g'
}

# fetch_statement <n>: print the problem statement as LaTeX, converted from the
# bare HTML that projecteuler.net/minimal=<n> serves. Its math is already TeX.
# Images and tables are dropped, so check those problems by hand.
fetch_statement() {
	curl -fsS --max-time 10 "https://projecteuler.net/minimal=$1" | sed -E \
		-e 's#^\$\$(.*)\$\$$#\\[\n  \1\n\\]#' \
		-e 's#<p>##g; s#</p>#\n#g' \
		-e 's#<br ?/?>#\n#g' \
		-e 's#<[^>]+>##g' \
		-e 's#\.\.\.#\\cdots#g' \
		-e 's#&lt;#<#g; s#&gt;#>#g; s#&nbsp;#~#g; s#&times;#$\\times$#g; s#&amp;#\\\&#g'
}

# write_proof <n> <pkg>: create <pkg>/<pkg>_proof.tex from the template below.
write_proof() {
	local n=$1 pkg=$2 author title statement template
	local tex="$pkg/${pkg}_proof.tex"
	if [[ -e $tex ]]; then
		echo "error: $tex already exists" >&2
		exit 1
	fi
	author=$(git config user.name || true)
	author=${author:-TODO}
	title=$(fetch_title "$n" || true)
	statement=$(fetch_statement "$n" || true)
	if [[ -z $title || -z $statement ]]; then
		echo "warning: could not fetch problem $n from projecteuler.net; using TODO placeholders" >&2
		title=${title:-TODO}
		statement=${statement:-Problem statement goes here (TODO)}
	fi
	template=$(cat <<'EOF'
\documentclass{article}
\usepackage{mathtools} % loads amsmath
\usepackage{amssymb}
\usepackage{amsthm}
\usepackage{hyperref}
\usepackage{cleveref} % must come after hyperref

\newtheorem{theorem}{Theorem}
\newtheorem{lemma}[theorem]{Lemma}
\theoremstyle{definition}
\newtheorem{definition}[theorem]{Definition}

\title{Project Euler Problem @N@: @TITLE@}
\author{@AUTHOR@}
\date{\today}

\begin{document}
\maketitle

\section*{Problem}

@STATEMENT@

\section*{Proof}

\begin{lemma}\label{lem:TODO}
  TODO
\end{lemma}
\begin{proof}
  TODO
\end{proof}

\end{document}
EOF
	)
	# Quoted replacements are inserted literally (no & or \ expansion).
	template=${template//@N@/"$n"}
	template=${template//@AUTHOR@/"${author^}"}
	template=${template//@TITLE@/"$title"}
	template=${template//@STATEMENT@/"$statement"}
	printf '%s\n' "$template" >"$tex"
	echo "created $tex"
}

if [[ $1 != --sync ]]; then
	[[ $1 =~ ^[0-9]+$ ]] || usage
	n=$((10#$1))
	((n > 0)) || usage
	pkg=$(printf 'p%03d' "$n")

	if [[ -e $pkg ]]; then
		if ((proof)); then
			# Existing problem: just add the proof, nothing else to regenerate.
			write_proof "$n" "$pkg"
			exit 0
		fi
		echo "error: $pkg already exists" >&2
		exit 1
	fi
	mkdir "$pkg"

	cat >"$pkg/$pkg.go" <<EOF
package $pkg

func Solve(limit int) uint64 {
	// TODO: solve problem $n
	return 0
}
EOF

	cat >"$pkg/${pkg}_test.go" <<EOF
package $pkg

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		limit int
		want  uint64
	}{
		// TODO: add the example from the problem statement and the real answer.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Solve(tt.limit)
			if got != tt.want {
				t.Errorf("Solve(%d) = %d; want %d", tt.limit, got, tt.want)
			}
		})
	}
}

func BenchmarkSolve(b *testing.B) {
	b.Skip("TODO: set a real benchmark input for problem $n, then remove this skip")
	for b.Loop() {
		_ = Solve(0)
	}
}
EOF

	cat >"$pkg/run.go" <<EOF
package $pkg

import (
	"fmt"
	"strconv"

	"$module/registry"
)

func init() {
	registry.Register($n, func(input string) (string, error) {
		target, err := strconv.Atoi(input)
		if err != nil {
			return "", fmt.Errorf("invalid limit: %w", err)
		}
		return strconv.FormatUint(Solve(target), 10), nil
	})
}
EOF
	((proof)) && write_proof "$n" "$pkg"
	echo "created $pkg/"
fi

# Verify each package registers the number matching its directory name.
bad=0
dirs=()
for d in p[0-9][0-9][0-9]; do
	[[ -d $d ]] || continue
	dirs+=("$d")
	want=$((10#${d#p}))
	got=$(grep -oP 'registry\.Register\(\s*\K[0-9]+' "$d"/*.go 2>/dev/null | cut -d: -f2 || true)
	if [[ -z $got ]]; then
		echo "warning: $d never calls registry.Register" >&2
		bad=1
	elif [[ $got != "$want" ]]; then
		echo "warning: $d registers problem(s) $(echo $got) but should register $want" >&2
		bad=1
	fi
done

{
	echo "package main"
	echo
	echo "// Each problem package registers itself in its init(); import new ones here."
	echo "// Generated by newprob.sh from the pNNN directories."
	echo "import ("
	for d in "${dirs[@]}"; do
		printf '\t_ "%s/%s"\n' "$module" "$d"
	done
	echo ")"
} >problems.go

gofmt -l . | grep -q . && gofmt -w .
go build ./...
echo "problems.go registers ${#dirs[@]} problem(s)"
exit $bad
