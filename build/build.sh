#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "$0")"  && pwd)"
WORK_DIR="$(dirname $SCRIPT_DIR)"

DIST_DIR=${WORK_DIR}/dist
WIN_APP_PATH=${DIST_DIR}/k8sproxy.exe
MAC_APP_PATH=${DIST_DIR}/k8sproxy
MAIN_PATH=${WORK_DIR}/main.go



# 编译打包window
function build_win() {
    echo '===开始构建Windows系统==='
    CGO_ENABLED=0 GOOS=windows go build -a -o ${WIN_APP_PATH}  ${MAIN_PATH}
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
  CGO_ENABLED=0 GOOS=darwin go build -a -o ${MAC_APP_PATH} ${MAIN_PATH}
  echo '===构建Mac系统结束==='
}

build_win
build_mac
