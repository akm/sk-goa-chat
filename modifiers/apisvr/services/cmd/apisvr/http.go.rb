# coding: utf-8

$content.sub!('"net/http"'){|key| ['"apisvr/lib/http/cors"', "\t"+ key].join("\n") }
$content.sub!('handler = httpmdlwr.Log(adapter)(handler)') do |s|
    ['handler = cors.NewFromEnv("APP_CORS_ALLOW_ORIGINS").Handle(handler)', s].join("\n")
end
