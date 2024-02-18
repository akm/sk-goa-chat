# coding: utf-8

pattern_path = ARGV.shift
if pattern_path.nil? || pattern_path.empty?
    puts "usage: ruby #{__FILE__} <pattern_path>"
    exit(1)
end

root_path = File.expand_path('../..', __FILE__)
parent_path = File.dirname(pattern_path)
target_pattern = File.join(root_path, parent_path, '**/*.go')

target_paths = Dir.glob(target_pattern)
target_paths.each do |target_path|
    $target_path = target_path
    $content = File.read(target_path)
    load(pattern_path) # パターンを読み込む
    open(target_path, 'w'){|f| f.puts($content) }
    system("go fmt #{target_path}", exception: true)
end
