defmodule Main do
  def stdin_values, do: IO.read(:eof) |> String.split()

  def parse_input([d]), do: d

  def solve("N"), do: "S"
  def solve("E"), do: "W"
  def solve("W"), do: "E"
  def solve("S"), do: "N"
  def solve("NE"), do: "SW"
  def solve("NW"), do: "SE"
  def solve("SE"), do: "NW"
  def solve("SW"), do: "NE"

  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end
end
