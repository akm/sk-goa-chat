# coding: utf-8

base_path = ARGV.shift
if base_path.nil? || base_path.empty?
    puts "usage: ruby #{__FILE__} <path_from_repository_root>"
    exit(1)
end

target_path = File.expand_path('../' + base_path, __FILE__)
pattern_path = File.expand_path('./' + base_path + '.rb', __FILE__)

$content = File.read(target_path)

load(pattern_path) # パターンを読み込む

open(target_path, 'w'){|f| f.puts($content) }

system("go fmt #{target_path}", exception: true)
