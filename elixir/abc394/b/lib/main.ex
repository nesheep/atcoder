defmodule Main do
  def stdin_values, do: IO.read(:eof) |> String.split()

  def parse_input([_ | t]), do: t

  def sort(l), do: Enum.sort(l, &(String.length(&1) <= String.length(&2)))

  def main do
    stdin_values()
    |> parse_input()
    |> sort()
    |> Enum.join()
    |> IO.puts()
  end
end
