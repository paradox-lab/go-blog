*******************************
chromedp
*******************************

保存Cookies
===============================

.. code-block:: go

    // GetCookies 将Cookies保存在本地
    func GetCookies(filename string) chromedp.Tasks {
        return chromedp.Tasks{
            chromedp.ActionFunc(func(ctx context.Context) error {
                cookies, err := network.GetAllCookies().Do(ctx)
                if err != nil {
                    return err
                }

                buf, err := json.Marshal(cookies)
                if err != nil {
                    fmt.Println("error:", err)
                }
                file,err:=os.OpenFile(filename,os.O_RDWR|os.O_CREATE,os.ModePerm)
                defer file.Close()
                if err !=nil{
                    fmt.Println("打开文件异常",err.Error())
                }else{
                    //【write】写入[]byte类型数据
                    n,err:=file.Write(buf)
                    if err!=nil{
                        fmt.Println("写入文件异常",err.Error())
                    }else{
                        fmt.Println("写入Ok：",n)
                    }
                }

                return nil
            }),
        }
    }

设置Cookies
=============================================

.. code-block:: go

    // AddCookies 读取Cookies文件，用于免登录打开页面
    func AddCookies(filename string) chromedp.Tasks{
        return chromedp.Tasks{
            chromedp.ActionFunc(func(ctx context.Context) error {
                // add cookies to chrome
                file, err := os.Open(filename)
                if err != nil {
                    fmt.Println("打开文件错误", err.Error())
                }

                defer file.Close()
                jsonBlob, _:=ioutil.ReadAll(file)
                var cookies []*network.CookieParam

                err = json.Unmarshal(jsonBlob, &cookies)
                if err != nil {
                    fmt.Println("error:", err)
                }
                err = network.ClearBrowserCookies().Do(ctx)
                if err!=nil{
                    return err
                }
                err = network.SetCookies(cookies).Do(ctx)
                if err!=nil{
                    return err
                }
                fmt.Println("OK")
                return nil
            }),
        }
    }