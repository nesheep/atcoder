defmodule Main do
  def stdin_values, do: IO.read(:eof) |> String.split()

  def parse_input([n]), do: String.to_integer(n)

  def solve(n) do
    1..n
    |> Enum.map(fn x ->
      1..n
      |> Enum.map(fn y ->
        if rem(Enum.min([x, y, n + 1 - x, n + 1 - y]), 2) == 0, do: ".", else: "#"
      end)
    end)
  end

  def output(a) do
    Enum.map(a, &Enum.join(&1, ""))
    |> Enum.join("\n")
  end

  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> output()
    |> IO.puts()
  end
end
