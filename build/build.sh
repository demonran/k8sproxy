#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "$0")"  && pwd)"
WORK_DIR="$(dirname $SCRIPT_DIR)"

echo "work dir is: ${WORK_DIR}"

DIST_DIR=${WORK_DIR}/dist
WIN_APP_PATH=${DIST_DIR}/k8sproxy.exe
MAC_APP_PATH=${DIST_DIR}/k8sproxy
CLIENT_MAIN_PATH=${WORK_DIR}/cmd/client/main.go
SERVER_MAIN_PATH=${WORK_DIR}/cmd/server/main.go
SERVER_APP_PATH=${DIST_DIR}/server



# 编译打包window
function build_win() {
    echo '===开始构建Windows系统==='
    CGO_ENABLED=0 GOOS=windows go build -a -o ${WIN_APP_PATH}  ${CLIENT_MAIN_PATH}
    # 打包
    cp ${SCRIPT_DIR}/run.bat ${DIST_DIR}
    cp ${SCRIPT_DIR}/wintun.dll ${DIST_DIR}
    cd ${DIST_DIR}
    zip -m ${DIST_DIR}/proxy.zip k8sproxy.exe run.bat wintun.dll
    echo '===构建Windows系统结束==='
}


# 编译mac
function build_mac() {
  echo '===开始构建Mac系统==='
  CGO_ENABLED=0 GOOS=darwin go build -a -o ${MAC_APP_PATH} ${CLIENT_MAIN_PATH}
  echo '===构建Mac系统结束==='
}

function build_server() {
  echo '===开始构建Server系统==='
  CGO_ENABLED=0 GOOS=linux go build -a -o ${SERVER_APP_PATH} ${SERVER_MAIN_PATH}
  echo '===构建Server系统结束==='
}

# 通过参数实现按需打包， -w 表示windows系统， -m 表示mac系统 -s 表示server
# 解析参数
while getopts "wms" opt; do
  case $opt in
    w)
      build_win
      ;;
    m)
      build_mac
      ;;
    s)
      build_server
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done

