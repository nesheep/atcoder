defmodule Main do
  def main do
    s = IO.read(:line) |> String.trim()
    ans = if String.match?(s, ~r/R.*M/), do: "Yes", else: "No"
    IO.puts(ans)
  end
end
