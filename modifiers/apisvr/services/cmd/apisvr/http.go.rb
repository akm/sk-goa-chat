# coding: utf-8

$content.sub!('"net/http"'){|key| ['"svrlib/http/cors"', "\t"+ key].join("\n") }
$content.sub!('handler = httpmdlwr.Log(adapter)(handler)') do |s|
    [   
        'handler = cors.NewFromEnv("APP_CORS_ALLOW_ORIGINS").Tap(func(c *cors.Cors) {',
        '    c.AllowHeaders = []string{"Content-Type", "Cookie"}', 
        '    c.AllowCredentials = "true"', 
        '    c.Logger = logger.Logger', 
        '}).Handle(handler)',
        s,
    ].join("\n")
end
