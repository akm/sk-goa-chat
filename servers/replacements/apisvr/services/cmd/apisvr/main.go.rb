# coding: utf-8

require 'fukowl'

Fukowl.replace(__FILE__.sub('replacements/', '').sub(/\.rb$/, '')) do |content|
    content.sub!('"context"'){|key| [
        'goaendpoints "applib/goa/endpoints"',
        '"applib/goa/goaext"',
        '"applib/goa/goasql"',
        key
    ].join("\n") }
    content.sub!('logger = log.New("chatapi", false)'){ 'logger = log.New("chatapi", dbgF != nil && *dbgF)' }

    error_handler_def = <<EOS
            wappers := goaendpoints.Wrappers{
                goaendpoints.ErrorHandler(goaext.DefaultErrorHandler(logger)),
                goasql.ConnectionEndpointWrapper(logger.Logger),
            }
EOS

    endpoints_cnt = 0
    content.gsub!(/(\w+)Endpoints = (\w+)\.NewEndpoints\((\w+)\)/) do |m|
        r = "#{$1}Endpoints = goaendpoints.Wrap[*#{$2}.Endpoints](#{$2}.NewEndpoints(#{$3}), wappers.Wrap)"
        r = error_handler_def + r  if endpoints_cnt == 0
        endpoints_cnt += 1
        r
    end
end
