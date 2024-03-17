def replace_file(target_path)
    content = File.read(target_path)
    yield(content)
    open(target_path, 'w'){|f| f.puts(content) }

    case File.extname(target_path)
    when '.go'
        system("go fmt #{target_path}", exception: true)
    end
end
