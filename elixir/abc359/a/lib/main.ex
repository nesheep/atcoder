defmodule Main do
  def main do
    IO.read(:line)
    |> String.trim()
    |> String.to_integer()
    |> then(fn n ->
      Stream.repeatedly(fn -> IO.read(:line) |> String.trim() end)
      |> Stream.take(n)
      |> Stream.filter(&(&1 == "Takahashi"))
      |> Enum.count()
    end)
    |> IO.puts()
  end
end
