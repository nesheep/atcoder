defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input([s]), do: String.to_charlist(s)

  defp solve(s) do
    Enum.sort(s)
    |> Enum.dedup()
    |> Stream.concat(Stream.cycle([nil]))
    |> Stream.zip(?a..?z)
    |> Enum.find_value(fn {a, b} -> if a != b, do: [b] end)
    |> to_string()
  end
end
