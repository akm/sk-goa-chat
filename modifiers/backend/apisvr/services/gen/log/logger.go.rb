# coding: utf-8

$content.sub!('"github.com/rs/zerolog"'){ '"apisvr/lib/log"' }

$content.gsub!('zerolog.'){ 'log.' }

$content.sub!('logger := log.New(os.Stderr).With().Timestamp().Str("service", serviceName).Logger()') {
    'logger := *log.NewServiceLogger(os.Stderr, serviceName)'
}

$content.sub!('logger.Info().Fields(fields).Msgf("HTTP Request")') {
    'logger.Info().Fields(fields)'
}
