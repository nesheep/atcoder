defmodule Main do
  def main do
    [_, m] = read_ints()
    h = read_ints()

    {_, ans} =
      h
      |> Enum.with_index(1)
      |> Enum.reduce_while({m, 0}, fn {e, i}, {am, _} = acc ->
        if am >= e, do: {:cont, {am - e, i}}, else: {:halt, acc}
      end)

    IO.puts(ans)
  end

  def read_ints() do
    IO.read(:line)
    |> String.trim()
    |> String.split(" ")
    |> Enum.map(&String.to_integer/1)
  end
end
