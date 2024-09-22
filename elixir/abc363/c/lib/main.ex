defmodule Main do
  def main do
    [n, k] = read_ints()
    s = IO.read(:line) |> String.trim()

    ans =
      String.graphemes(s)
      |> permutations()
      |> Enum.uniq()
      |> Enum.map(&Enum.join(&1))
      |> Enum.map(fn t -> 0..(n - k) |> Enum.map(&String.slice(t, &1, k)) end)
      |> Enum.map(fn l -> l |> Enum.map(&(&1 != String.reverse(&1))) end)
      |> Enum.map(&Enum.all?(&1))
      |> Enum.count(& &1)

    IO.puts(ans)
  end

  def read_ints() do
    IO.read(:line)
    |> String.trim()
    |> String.split(" ")
    |> Enum.map(&String.to_integer/1)
  end

  def permutations([]), do: [[]]
  def permutations(l), do: for(e <- l, p <- permutations(l -- [e]), do: [e | p])
end
