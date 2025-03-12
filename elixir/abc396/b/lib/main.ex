defmodule Main do
  def stdin_values, do: IO.read(:eof) |> String.split()

  def take_int([x | input]), do: {String.to_integer(x), input}

  def take_list(f, n) do
    fn input ->
      {l, input} =
        Stream.iterate({[], input}, fn {l, input} ->
          {v, input} = f.(input)
          {[v | l], input}
        end)
        |> Stream.take(n + 1)
        |> Enum.reduce(nil, fn elem, _ -> elem end)

      {Enum.reverse(l), input}
    end
  end

  def take_query(["1", x | input]), do: {{:query1, String.to_integer(x)}, input}
  def take_query(["2" | input]), do: {:query2, input}

  def parse_input(input) do
    {q, input} = take_int(input)
    {l, _} = take_list(&take_query/1, q).(input)
    l
  end

  def run_query({:query1, x}, {acc, _}), do: {[x | acc], nil}
  def run_query(:query2, {[h | t], _}), do: {t, h}
  def run_query(:query2, {[], _}), do: {[], 0}

  def solve(queries) do
    queries
    |> Stream.scan({[], nil}, &run_query/2)
    |> Stream.map(fn {_, x} -> x end)
    |> Stream.filter(&(&1 != nil))
    |> Enum.to_list()
  end

  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> Enum.each(&IO.puts/1)
  end
end
