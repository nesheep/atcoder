defmodule Main do
  def main do
    s = IO.read(:line) |> String.trim()
    n = s |> String.to_charlist() |> Enum.filter(&(&1 in ?a..?z)) |> Enum.count()
    f = if n * 2 >= String.length(s), do: &String.downcase/1, else: &String.upcase/1
    IO.puts(s |> f.())
  end
end
