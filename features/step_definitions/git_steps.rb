Given "the git repository" do
    run_simple("git init")
    run_simple("git commit --allow-empty -m \"Add tmp\"")
end

Given "the git branch \"$branch\"" do |branch|
    run_simple("git checkout -b " + branch)
    run_simple("git checkout master")
end

Given "I run `git checkout $branch`" do |branch|
    run_simple("git checkout " + branch)
end
