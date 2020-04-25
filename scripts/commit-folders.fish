#!/usr/bin/env fish

function __list_directories \
    --description "List LeetCode problems folders" \
    --argument-names path
    ls $path | grep --color=never "^\d\{4\}" | grep --invert-match "^0000"
end

function __problem_url \
    --description "Get problem url" \
    --argument-names problem_dir
    set --local problem_name (string sub --start 6 $problem_dir)
    set --local problem_url \
        (printf "https://leetcode.com/problems/%s/" (string replace --all _ - $problem_name))
    echo $problem_url
end

function __check_page \
    --argument-names url
    command curl --silent \
        --head \
        --output /dev/null \
        --write-out "%{http_code}" \
        $url
end

function __fetch_readme \
    --argument-names url

    # set --local code 201
    set --local code (command curl --silent \
          # --head \
          --output .tmp/problem.html \
          --write-out "%{http_code}" \
          $url)

    if test $code -eq 200
        echo "YAYA"
    else
        echo "Non"
    end
end

function commit_problems \
    --description "Commit LeetCode problems"
    for dir in (__list_directories .)
        set --local url (__problem_url $dir)
        echo (__fetch_readme $url)
        return
    end
end

function graphql
    set -l res (command curl 'https://leetcode.com/graphql' \
        -H 'authority: leetcode.com' \
        -H 'accept: */*' \
        -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.113 Safari/537.36' \
        -H 'content-type: application/json' \
        -H 'origin: https://leetcode.com' \
        -H 'referer: https://leetcode.com/problems/move-zeroes/' \
        -H 'accept-language: en-US,en;q=0.9,ru;q=0.8,nl;q=0.7' \
        --data-binary '{"operationName":"questionData","variables":{"titleSlug":"move-zeroes"},"query":"query questionData($titleSlug: String\u0021) {\\n  question(titleSlug: $titleSlug) {\\n    questionId\\n    questionFrontendId\\n    boundTopicId\\n    title\\n    titleSlug\\n content\\n  difficulty\\n }\\n}\\n"}' \
        --compressed)

    printf $res
end

# commit_problems
graphql
