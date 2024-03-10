require 'fukowl'

Dir.glob(File.expand_path('../*.go', __FILE__.sub('replacements/', ''))) do |file|
    Fukowl.replace(file) do |content|
        content.sub!('"time"'){|key| '"applib/time"' }
    end
end
