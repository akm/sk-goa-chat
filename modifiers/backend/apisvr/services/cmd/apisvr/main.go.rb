# coding: utf-8

$content.sub!('chatapi "apisvr/services"'){|key| ['"apisvr/lib/goa/goaext"', "\t"+ key].join("\n") }
$content.sub!('logger = log.New("chatapi", false)'){ 'logger = log.New("chatapi", dbgF != nil && *dbgF)' }

error_handler_def = <<EOS
        var eh func(err error) error
        if os.Getenv("STAGE") != "local" {
            eh = goaext.LoggerErrorHandlerFunc(logger)
        } else {
            eh = goaext.StderrErrorHandler
        }
EOS
endpoints_cnt = 0
$content.gsub!(/(\w+)Endpoints = (\w+)\.NewEndpoints\((\w+)\)/) do |m|
    r = "#{$1}Endpoints = goaext.ErrorHandledEndpoints[*#{$2}.Endpoints](#{$2}.NewEndpoints(#{$3}), eh)"
    r = error_handler_def + r  if endpoints_cnt == 0
    endpoints_cnt += 1
    r
end
