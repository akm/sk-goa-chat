# coding: utf-8

require './replace_file'

replace_file(__FILE__.sub('replacements/', '').sub(/\.rb$/, '')) do |content|
    content.sub!('"github.com/rs/zerolog"'){ '"applib/log"' }

    content.gsub!('zerolog.'){ 'log.' }

    content.sub!('logger := log.New(os.Stderr).With().Timestamp().Str("service", serviceName).Logger()') {
        'logger := *log.NewServiceLogger(os.Stderr, serviceName)'
    }

    content.sub!('logger.Info().Fields(fields).Msgf("HTTP Request")') {
        'logger.Info().Fields(fields)'
    }
end
