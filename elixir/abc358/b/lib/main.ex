defmodule Main do
  def main do
    [_, a] = read_ints()
    ts = read_ints()
    solve(ts, a, []) |> Enum.each(&IO.puts/1)
  end

  def read_ints() do
    IO.read(:line)
    |> String.trim()
    |> String.split(" ")
    |> Enum.map(&String.to_integer/1)
  end

  def solve([], _, acc), do: acc |> Enum.reverse()
  def solve([h | t], a, []), do: solve(t, a, [h + a])
  def solve([h | t], a, [ph | pt]) when h < ph, do: solve(t, a, [ph + a, ph | pt])
  def solve([h | t], a, acc), do: solve(t, a, [h + a | acc])
end
