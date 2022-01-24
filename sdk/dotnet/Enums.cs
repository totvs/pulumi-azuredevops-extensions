// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureDevopsExtensions
{
    [EnumType]
    public readonly struct RoleName : IEquatable<RoleName>
    {
        private readonly string _value;

        private RoleName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static RoleName Reader { get; } = new RoleName("Reader");
        public static RoleName Administrator { get; } = new RoleName("Administrator");
        public static RoleName User { get; } = new RoleName("User");

        public static bool operator ==(RoleName left, RoleName right) => left.Equals(right);
        public static bool operator !=(RoleName left, RoleName right) => !left.Equals(right);

        public static explicit operator string(RoleName value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RoleName other && Equals(other);
        public bool Equals(RoleName other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ScopeName : IEquatable<ScopeName>
    {
        private readonly string _value;

        private ScopeName(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScopeName VariableGroup { get; } = new ScopeName("VariableGroup");
        public static ScopeName ServiceEndpoint { get; } = new ScopeName("ServiceEndpoint");
        public static ScopeName Environment { get; } = new ScopeName("Environment");

        public static bool operator ==(ScopeName left, ScopeName right) => left.Equals(right);
        public static bool operator !=(ScopeName left, ScopeName right) => !left.Equals(right);

        public static explicit operator string(ScopeName value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScopeName other && Equals(other);
        public bool Equals(ScopeName other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
