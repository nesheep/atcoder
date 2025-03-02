defmodule Main do
  def stdin_values, do: IO.read(:eof) |> String.split()

  def parse_input([_ | t]), do: Enum.map(t, &String.to_integer/1)

  def solve(a) when is_list(a) do
    Enum.chunk_every(a, 2, 1, :discard)
    |> Enum.all?(fn [x, y] -> x < y end)
  end

  def output(true), do: "Yes"
  def output(false), do: "No"

  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> output()
    |> IO.puts()
  end
end
