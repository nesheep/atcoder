defmodule Main do
  def stdin_values, do: IO.read(:eof) |> String.split()

  def solve([a1, a2, a3]) when a1 * a2 == a3, do: true
  def solve([a2, a3, a1]) when a1 * a2 == a3, do: true
  def solve([a3, a1, a2]) when a1 * a2 == a3, do: true
  def solve(_), do: false

  def output(true), do: "Yes"
  def output(false), do: "No"

  def main do
    stdin_values()
    |> Enum.map(&String.to_integer/1)
    |> solve()
    |> output()
    |> IO.puts()
  end
end
