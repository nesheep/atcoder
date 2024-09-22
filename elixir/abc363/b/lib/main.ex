defmodule Main do
  def main do
    [_, t, p] = read_ints()
    ls = read_ints()

    ans =
      ls
      |> Enum.sort(:desc)
      |> Enum.at(p - 1)
      |> then(&(t - &1))
      |> max(0)

    IO.puts(ans)
  end

  def read_ints() do
    IO.read(:line)
    |> String.trim()
    |> String.split(" ")
    |> Enum.map(&String.to_integer/1)
  end
end
