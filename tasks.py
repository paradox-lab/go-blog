from invoke import task


@task
def go_test(c):
    """执行go测试

    完整命令: go test -json ./tests/... | go-test-report -o report/html/report-$(date "+%Y%m%d%H%M%S").html
    """
    c.run("go test -json ./tests/...")