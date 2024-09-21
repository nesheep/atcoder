defmodule Main do
  def main do
    [a, b, c, d, e, f] = read_ints()
    [g, h, i, j, k, l] = read_ints()
    ans = a < j and b < k and c < l and d > g and e > h and f > i
    IO.puts(if ans, do: "Yes", else: "No")
  end

  def read_ints() do
    IO.read(:line)
    |> String.trim()
    |> String.split(" ")
    |> Enum.map(&String.to_integer/1)
  end
end
