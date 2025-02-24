defmodule Main do
  def stdin_values, do: IO.read(:eof) |> String.split()

  def take_int([h | t]), do: {String.to_integer(h), t}

  def parse_input(input) do
    {n, input} = take_int(input)
    {t, input} = take_int(input)
    {a, _} = take_int(input)
    {n, t, a}
  end

  def solve({n, t, a}), do: n < t * 2 || n < a * 2

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
