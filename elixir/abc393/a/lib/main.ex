defmodule Main do
  def stdin_values, do: IO.read(:eof) |> String.split()

  def parse_input([s1, s2]), do: {s1, s2}

  def solve({"sick", "sick"}), do: 1
  def solve({"sick", "fine"}), do: 2
  def solve({"fine", "sick"}), do: 3
  def solve({"fine", "fine"}), do: 4

  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end
end
