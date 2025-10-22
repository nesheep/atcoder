defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input(x), do: Enum.map(x, &String.to_integer/1)

  defp solve([s, a, b, x]) do
    Step.new(s, a, b)
    |> Stream.take(x)
    |> Enum.sum()
  end
end

defmodule Step do
  defstruct s: 0, a: 0, b: 0

  def new(s, a, b), do: %Step{s: s, a: a, b: b}
end

defimpl Enumerable, for: Step do
  def count(_step), do: {:error, __MODULE__}
  def member?(_step, _value), do: {:error, __MODULE__}
  def slice(_step), do: {:error, __MODULE__}
  def reduce(step, acc, fun), do: reduce(step, acc, fun, 0)

  defp reduce(_stem, {:halt, acc}, _fun, _i), do: {:halted, acc}
  defp reduce(step, {:suspend, acc}, fun, i), do: {:suspended, acc, &reduce(step, &1, fun, i)}

  defp reduce(%Step{s: s, a: a, b: b} = step, {:cont, acc}, fun, i) do
    t = if rem(i, a + b) < a, do: s, else: 0
    reduce(step, fun.(t, acc), fun, i + 1)
  end
end
