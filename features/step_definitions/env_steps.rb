Given "the environment" do
    run_simple("zenv local")
end

Given "the not-recursive environment" do
    run_simple("zenv local --not-recursive", false)
end

Given "the exclusive environment" do
    run_simple("zenv local --exclusive", false)
end

Given "the not-recursive and exclusive environment" do
    run_simple("zenv local --not-recursive --exclusive", false)
end

Given "the global environment" do
    run_simple("zenv global", false)
end

Given "the not-recursive global environment" do
    run_simple("zenv global --not-recursive", false)
end

Given "the exclusive global environment" do
    run_simple("zenv global --exclusive", false)
end

Given "the not-recursive and exclusive global environment" do
    run_simple("zenv global --not-recursive --exclusive", false)
end
