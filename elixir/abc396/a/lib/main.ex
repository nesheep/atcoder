defmodule Main do
  def stdin_values, do: IO.read(:eof) |> String.split()

  def parse_input([_ | a]), do: Enum.map(a, &String.to_integer/1)

  def solve(a) when is_list(a) do
    Enum.chunk_every(a, 3, 1, :discard)
    |> Enum.any?(fn v -> MapSet.new(v) |> MapSet.size() == 1 end)
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
