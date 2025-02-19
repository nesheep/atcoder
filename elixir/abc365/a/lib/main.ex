defmodule Main do
  def main do
    y = IO.read(:line) |> String.trim() |> String.to_integer()
    {:ok, d} = Date.new(y, 12, 31)
    ans = Date.day_of_year(d)
    IO.puts(ans)
  end
end
