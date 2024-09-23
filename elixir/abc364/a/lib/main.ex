defmodule Main do
  def main do
    n =
      IO.read(:line)
      |> String.trim()
      |> String.to_integer()

    ans =
      Stream.repeatedly(fn -> IO.read(:line) |> String.trim() end)
      |> Stream.take(n - 1)
      |> Stream.chunk_every(2, 1, :discard)
      |> Stream.map(fn [a, b] -> a == "sweet" && b == "sweet" end)
      |> Enum.any?()

    IO.puts(if ans, do: "No", else: "Yes")
  end
end
