defmodule Main do
  def main do
    n = IO.read(:line) |> String.trim() |> String.to_integer()
    ls = IO.read(:line) |> String.trim() |> String.split(" ") |> Enum.map(&String.to_integer/1)

    ans =
      1..n
      |> Enum.map(fn i ->
        Enum.with_index(ls)
        |> Enum.filter(fn {a, _} -> a == i end)
        |> Enum.map(fn {_, j} -> j + 1 end)
      end)
      |> Enum.filter(fn a -> Enum.max(a) - Enum.min(a) == 2 end)
      |> Enum.count()

    IO.puts(ans)
  end
end
