defmodule Main do
  def main do
    IO.read(:line)
    |> String.trim()
    |> String.to_integer()
    |> then(&(100 - rem(&1, 100)))
    |> IO.puts()
  end
end
