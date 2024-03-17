require './replace_file'

Dir.glob(File.expand_path('../*.go', __FILE__.sub('replacements/', ''))) do |file|
    replace_file(file) do |content|
        content.sub!('"time"'){|key| '"applib/time"' }
    end
end
